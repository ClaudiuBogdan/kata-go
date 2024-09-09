#!/bin/bash

PRACTICE_DIR="practice"

TEMPLATES_DIR="templates"

# Function to generate a new day
generate_new_day() {
    local latest_day=$(ls -1 $PRACTICE_DIR | grep -E '^day_[0-9]+$' | sort -V | tail -n 1)
    if [ -z "$latest_day" ]; then
        new_day="day_1"
    else
        local day_num=$(echo $latest_day | cut -d'_' -f2)
        new_day="day_$((day_num + 1))"
    fi
    
    mkdir -p "$PRACTICE_DIR/$new_day"
    echo "Created new practice day: $new_day"
}

# Function to copy a kata to a specific day
copy_kata() {
    local template_path="$TEMPLATES_DIR/$1"
    local target_day=$2

    if [ ! -d "$template_path" ]; then
        echo "Error: Template path does not exist: $template_path"
        exit 1
    fi

    if [ -z "$target_day" ]; then
        target_day=$(ls -1 $PRACTICE_DIR | grep -E '^day_[0-9]+$' | sort -V | tail -n 1)
        if [ -z "$target_day" ]; then
            echo "Error: No practice days found. Please generate a new day first."
            exit 1
        fi
    elif [ ! -d "$PRACTICE_DIR/$target_day" ]; then
        echo "Error: Practice day does not exist: $target_day"
        exit 1
    fi

    local kata_name=$(basename $template_path)
    local target_path="$PRACTICE_DIR/$target_day/$kata_name"

    if [ -d "$target_path" ]; then
        echo "Error: Kata '$kata_name' already exists in $target_day"
        exit 1
    fi

    cp -R "$template_path" "$target_path"
    echo "Copied kata '$kata_name' to $target_day"
}

# Main script logic
case "$1" in
    new-day)
        generate_new_day
        ;;
    copy-kata)
        if [ "$#" -lt 2 ]; then
            echo "Usage: $0 copy-kata <template_path> [target_day]"
            exit 1
        fi
        copy_kata "$2" "$3"
        ;;
    *)
        echo "Usage: $0 {new-day|copy-kata <template_path> [target_day]}"
        exit 1
        ;;
esac