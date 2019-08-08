# Copyright (c) 2019, Microsoft Corporation, Sean Hinchee
# Licensed under the MIT License.

# Convert a .vmdk VMWare file to .vhdx Hyper-V file
# usage: vmdk2vhdx file.vmdk file.vhdx

param (
	[string]$from,
	[string]$to
)

function usage {
	$0 = split-path $MyInvocation.PSCommandPath -Leaf
	Write-Output "usage: $0 file.vmdk file.vhdx"
	exit "usage"
}

if(!$from -Or !$to) {
	usage
}

# As per: https://blogs.msdn.microsoft.com/timomta/2015/06/11/how-to-convert-a-vmware-vmdk-to-hyper-v-vhd/
Import-Module 'C:\Program Files\Microsoft Virtual Machine Converter\MvmcCmdlet.psd1'

ConvertTo-MvmcVirtualHardDisk -SourceLiteralPath $p1 -VhdType DynamicHardDisk -VhdFormat vhdx -destination $p2
