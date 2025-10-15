$ErrorActionPreference = "Stop"

# Where fonts will be installed (user-only)
$FontsDir = "$env:LOCALAPPDATA\Microsoft\Windows\Fonts"
New-Item -ItemType Directory -Path $FontsDir -Force | Out-Null

# URLs for font zips
$JetBrainsUrl = "https://download.jetbrains.com/fonts/JetBrainsMono-2.304.zip"
$FiraCodeUrl = "https://github.com/tonsky/FiraCode/releases/latest/download/Fira_Code_v6.2.zip"

# Temp download locations
$Temp = "$env:TEMP\vscode_fonts"
Remove-Item $Temp -Recurse -Force -ErrorAction SilentlyContinue
New-Item -ItemType Directory -Path $Temp | Out-Null

# Download and extract
Invoke-WebRequest -Uri $JetBrainsUrl -OutFile "$Temp\jetbrains.zip"
Invoke-WebRequest -Uri $FiraCodeUrl -OutFile "$Temp\firacode.zip"

Expand-Archive -Path "$Temp\jetbrains.zip" -DestinationPath "$Temp\jetbrains"
Expand-Archive -Path "$Temp\firacode.zip" -DestinationPath "$Temp\firacode"

# Copy .ttf files
Get-ChildItem "$Temp\jetbrains\**\*.ttf" -Recurse | Copy-Item -Destination $FontsDir -Force

# Register fonts (user scope)
$ShellApp = New-Object -ComObject Shell.Application
$FontFolder = $ShellApp.Namespace(0x14)

Get-ChildItem $FontsDir -Filter "*.ttf" | ForEach-Object {
    $FontFolder.CopyHere($_.FullName)
}

Write-Host "✅ Fonts installed: JetBrains Mono & Fira Code"
