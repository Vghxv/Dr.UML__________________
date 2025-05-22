# Save the current working directory
$OriginalLocation = Get-Location

# Get the directory of the current script
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
$env:APP_ROOT = $ScriptDir

# Change to script directory
Set-Location -Path $ScriptDir

# Start Wails development server
Write-Host "Starting Wails development server..."
wails dev

# Return to original location
Set-Location -Path $OriginalLocation
