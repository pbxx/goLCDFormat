# This script converts the SVG mockup image to PNG for use in README.md

param(
    [Parameter(Mandatory = $true)]
    [string]$Name,
    [switch]$Line
)

$MockupDir = ".\mockups"

# Create the dir if it doesnt exist
if (!(Test-Path -Path $MockupDir)) {
    New-Item -ItemType Directory -Path $MockupDir > $null
}
if ($Line) {
    $File = "./linemockup.svg"
} else {
    $File = "./mockscreen.svg"
}


inkscape $File -o "./mockups/$Name.png"

Write-Host "Finished! Mockup image saved at assets/mockups/$Name.png"