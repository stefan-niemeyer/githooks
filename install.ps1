#Requires -Version 5.1
<#
.SYNOPSIS
    Install githooks on Windows.
.DESCRIPTION
    Downloads the latest githooks release for Windows and extracts it
    to the current directory.
.EXAMPLE
    irm https://raw.githubusercontent.com/stefan-niemeyer/githooks/main/install.ps1 | iex
#>

$ErrorActionPreference = 'Stop'

$ARCH = ""
$DOWNLOAD_URL = ""

function Get-WindowsArch {
    $arch = $env:PROCESSOR_ARCHITECTURE

    if ([string]::IsNullOrWhiteSpace($arch)) {
        throw "PROCESSOR_ARCHITECTURE is empty"
    }

    $arch = $arch.ToLower()

    if ($arch -eq 'x86' -and -not [string]::IsNullOrWhiteSpace($env:PROCESSOR_ARCHITECTURE)) {
        $arch = $env:PROCESSOR_ARCHITECTURE.ToLower()
    }

    switch ($arch) {
        'amd64' { return 'amd64' }
        'x64'   { return 'amd64' }
        'arm64' { return 'arm64' }
        default { throw "$arch isn't supported" }
    }
}

function Get-LatestDownloadUrl {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Arch
    )

    $apiUrl = 'https://api.github.com/repos/stefan-niemeyer/githooks/releases/latest'
    $assetName = "windows-$Arch.zip"
    $headers = @{
        'User-Agent' = 'githooks-installer-powershell'
        'Accept'     = 'application/vnd.github+json'
    }

    try {
        $release = Invoke-RestMethod -Uri $apiUrl -Headers $headers -UseBasicParsing
    }
    catch {
        throw "Failed to query GitHub release API: $($_.Exception.Message)"
    }

    if (-not $release -or -not $release.assets) {
        throw "GitHub release response did not contain any assets"
    }

    $asset = $release.assets |
        Where-Object { $_.name -eq $assetName } |
        Select-Object -First 1

    if (-not $asset) {
        $asset = $release.assets |
            Where-Object { $_.browser_download_url -like "*$assetName" } |
            Select-Object -First 1
    }

    if (-not $asset -or [string]::IsNullOrWhiteSpace($asset.browser_download_url)) {
        throw "No release asset found for $assetName"
    }

    return $asset.browser_download_url
}

function Expand-ZipFile {
    param(
        [Parameter(Mandatory = $true)]
        [string]$ZipPath,

        [Parameter(Mandatory = $true)]
        [string]$DestinationPath
    )

    try {
        Expand-Archive -LiteralPath $ZipPath -DestinationPath $DestinationPath -Force
        return
    }
    catch {
    }

    try {
        Add-Type -AssemblyName System.IO.Compression.FileSystem
        [System.IO.Compression.ZipFile]::ExtractToDirectory($ZipPath, $DestinationPath)
        return
    }
    catch {
        throw "Unpacking failed: $($_.Exception.Message)"
    }
}

try {
    if ([string]::IsNullOrWhiteSpace($ARCH)) {
        $ARCH = Get-WindowsArch
    }

    if ([string]::IsNullOrWhiteSpace($DOWNLOAD_URL)) {
        $DOWNLOAD_URL = Get-LatestDownloadUrl -Arch $ARCH
    }

    $filename = Split-Path -Path $DOWNLOAD_URL -Leaf
    $zipPath = Join-Path -Path $PWD -ChildPath $filename

    Write-Host "Downloading githooks from $DOWNLOAD_URL ..."

    $webRequestParams = @{
        Uri             = $DOWNLOAD_URL
        OutFile         = $zipPath
        Headers         = @{ 'User-Agent' = 'githooks-installer-powershell' }
        UseBasicParsing = $true
    }

    try {
        Invoke-WebRequest @webRequestParams
    }
    catch {
        throw "Failed to download $filename : $($_.Exception.Message)"
    }

    Expand-ZipFile -ZipPath $zipPath -DestinationPath $PWD

    Write-Host "Installation complete! Please copy githooks to a folder in your PATH"
}
catch {
    Write-Host ""
    Write-Host $_.Exception.Message
    Write-Host ""
    exit 1
}
finally {
    if ($zipPath -and (Test-Path -LiteralPath $zipPath)) {
        Remove-Item -LiteralPath $zipPath -Force -ErrorAction SilentlyContinue
    }
}