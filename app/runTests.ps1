# Save the current working directory
$OriginalLocation = Get-Location

# Get the directory of the current script
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
$env:APP_ROOT = $ScriptDir

# Change to script directory
Set-Location -Path $ScriptDir

# Run Go tests with coverage
Write-Host "Running tests with coverage..."
if (-not (go test -v -coverprofile="coverage.out" ./...)) {
    Write-Error "Tests failed."
    Set-Location -Path $OriginalLocation
    exit 1
}

# Generate coverage report
Write-Host "Generating coverage report..."
go tool cover -func="coverage.out"

# Return to original location
Set-Location -Path $OriginalLocation
