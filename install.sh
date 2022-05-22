#!/bin/sh

OSTYPE=""
ARCH=""
DOWNLOAD_URL=""

if [ "x${OSTYPE}" = "x" ]; then
  case $(uname) in
  "Linux")
    OSTYPE="linux"
    ;;
  "Darwin")
    OSTYPE="darwin"
    ;;
  *)
    echo 'Warning: Only Linux and MacOS operating systems are currently supported! For Windows-OS,
    please copy the tar.gz file directly.'
    exit 1
    ;;
  esac
fi

if [ "x${ARCH}" = "x" ]; then
  case "$(uname -m)" in
  x86_64)
    ARCH=amd64
    ;;
  armv8*)
    ARCH=arm64
    ;;
  aarch64*)
    ARCH=arm64
    ;;
  *)
    echo "${ARCH}, isn't supported"
    exit 1
    ;;
  esac
fi

if [ "x${DOWNLOAD_URL}" = "x" ]; then
  DOWNLOAD_URL="$(curl -sL "https://api.github.com/repos/xiabai84/githooks/releases/latest" |
    grep browser_download_url |
    cut -d '"' -f 4 |
    grep "$OSTYPE-$ARCH")"
fi

filename="${DOWNLOAD_URL##*/}"

echo "Downloading githooks from ${DOWNLOAD_URL} ..."

trap 'rm -f $filename' EXIT

curl -fsLO "$DOWNLOAD_URL"

if [ $? -ne 0 ]; then
  echo ""
  echo "Failed to download $filename!"
  echo ""
  echo "Please verify the version you are trying to download."
  echo ""
  exit
fi

ret='0'
command -v tar >/dev/null 2>&1 || { ret='1'; }
if [ "$ret" -eq 0 ]; then
  tar -xzf "${filename}"
  echo "Installation Complete!"
else
  echo "$filename Download Complete!"
  echo ""
  echo "Try to unpack the ${filename} failed."
  echo "tar: command not found, please unpack the ${filename} manually."
  exit 1
fi
