#!/bin/sh

OSTYPE=""
ARCH=""
DOWNLOAD_URL=""

if [ -z "${OSTYPE}" ]; then
  case $(uname) in
  "Linux")
    OSTYPE="linux"
    ;;
  "Darwin")
    OSTYPE="darwin"
    ;;
  *)
    echo -e "Warning: Only Linux and macOS operating systems are currently supported!\nFor Windows-OS,
    please copy the tar.gz file directly."
    exit 1
    ;;
  esac
fi

if [ -z "${ARCH}" ]; then
  case "$(uname -m)" in
  x86_64)
    ARCH=amd64
    ;;
  armv8*)
    ARCH=arm64
    ;;
  aarch64* | arm64*)
    ARCH=arm64
    ;;
  *)
    echo "$(uname -m), isn't supported"
    exit 1
    ;;
  esac
fi

if [ -z "${DOWNLOAD_URL}" ]; then
  DOWNLOAD_URL=$(curl -sL "https://api.github.com/repos/stefan-niemeyer/githooks/releases/latest" | \
    grep -Ee "browser_download_url.*$OSTYPE-$ARCH.tar.gz" | \
    sed -Ee 's/^ *"browser_download_url": *"(.*)"/\1/g')
fi

filename="${DOWNLOAD_URL##*/}"
echo "Downloading githooks from ${DOWNLOAD_URL} ..."

trap 'rm -f $filename' EXIT

if ! curl -fsLO "$DOWNLOAD_URL"; then
  echo -e "\nFailed to download $filename!\n\nPlease verify the version you are trying to download.\n"
  exit 1
fi

if command -v tar >/dev/null 2>&1
then
  tar -xzf "${filename}"
  echo "Installation Complete! Please copy githooks in a folder in your PATH"
else
  echo -e "$filename Download complete!\nUnpacking ${filename} failed."
  echo "tar: command not found, please unpack ${filename} manually."
  exit 1
fi
