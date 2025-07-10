#!/bin/bash

# Script to replace uint64 with int64 in manticore.yml and output to another file

# Check if manticore.yml exists
if [ ! -f "manticore.yml" ]; then
    echo "Error: manticore.yml file not found in current directory"
    exit 1
fi

# Define input and output files
input_file="manticore.yml"
output_file="manticore_int64.yml"

# Replace uint64 with int64 and save to output file
sed 's/uint64/int64/g' "$input_file" > "$output_file"

# Count the number of replacements made
original_count=$(grep -c "uint64" "$input_file" 2>/dev/null || echo "0")
new_count=$(grep -c "int64" "$output_file" 2>/dev/null || echo "0")

echo "Replacement completed successfully!"
echo "Input file: $input_file"
echo "Output file: $output_file"
echo "Original uint64 occurrences: $original_count"
echo "New int64 occurrences: $new_count"

# Verify the replacement worked
if [ "$original_count" -gt 0 ]; then
    echo "Verification: Checking if any uint64 remains in output file..."
    remaining_uint64=$(grep -c "uint64" "$output_file" 2>/dev/null || echo "0")
    if [ "$remaining_uint64" -eq 0 ]; then
        echo "✓ All uint64 values have been successfully replaced with int64"
    else
        echo "⚠ Warning: $remaining_uint64 uint64 values still remain in output file"
    fi
else
    echo "Note: No uint64 values were found in the original file"
fi 