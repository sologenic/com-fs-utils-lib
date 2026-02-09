#!/bin/bash
set -e

PROTO_DIR=$(dirname $(dirname "$0"))
cd $PROTO_DIR

# Function to get list of model directories (excluding bin and build)
get_model_dirs() {
    find . -maxdepth 1 -type d ! -name "." ! -name "bin" ! -name "build" ! -name "node_modules" ! -name ".git" | sed 's|^\./||' | sort
}

# Function to check if a directory contains proto files
has_proto_files() {
    local dir=$1
    [ -n "$(find "$dir" -maxdepth 1 -name "*.proto" 2>/dev/null)" ]
}

# Function to build a specific model
build_model() {
    local model_dir=$1
    
    if [ ! -d "$model_dir" ]; then
        echo "Error: Directory '$model_dir' does not exist"
        exit 1
    fi
    
    if ! has_proto_files "$model_dir"; then
        echo "Error: No .proto files found in '$model_dir'"
        exit 1
    fi
    
    echo "Building model: $model_dir"
    
    # Remove the specific model's build directory for a clean build
    if [ -d "build/$model_dir" ]; then
        echo "  Removing existing build directory for $model_dir..."
        rm -rf "build/$model_dir"
    fi
    
    # Generate Go files for this model
    for proto in $(find "$model_dir" -maxdepth 1 -name "*.proto"); do
        echo "  Generating Go files for: $proto"
        protoc \
            --proto_path=. \
            --proto_path="$(go env GOPATH)/src" \
            --go_out=. \
            --go_opt=paths=source_relative \
            "$proto"
    done
    
    # Generate TypeScript files
    if [ ! -d "node_modules" ]; then
        echo "  Installing npm dependencies..."
        npm i
    fi
    
    for proto in $(find "$model_dir" -maxdepth 1 -name "*.proto"); do
        echo "  Generating TypeScript files for: $proto"
        protoc \
            --plugin=protoc-gen-ts_proto="../node_modules/.bin/protoc-gen-ts_proto" \
            --ts_proto_out=. \
            --ts_proto_opt=esModuleInterop=true \
            --proto_path=. \
            --proto_path="$(go env GOPATH)/src" \
            "$proto"
    done
    
    # Build TypeScript
    echo "  Building TypeScript..."
    npm run build
    
    # Git add for this specific model
    if [ -d "build/$model_dir" ]; then
        git add "build/$model_dir/"
    fi
    
    # Add TypeScript files in the model directory
    for ts_file in $(find "$model_dir" -maxdepth 1 -name "*.ts" 2>/dev/null); do
        git add "$ts_file"
    done
    
    echo "✓ Build complete for $model_dir"
}

# Function to build all models
build_all() {
    echo "Building all models..."
    
    # Remove build directory for a clean build
    if [ -d "build" ]; then
        echo "  Removing existing build directory..."
        rm -rf build
    fi
    
    # Generate Go files for all proto files
    for proto in $(find . -name "*.proto" -not -path "../node_modules/*" -not -path "./build/*"); do
        echo "  Generating Go files for: $proto"
        protoc \
            --proto_path=. \
            --proto_path="$(go env GOPATH)/src" \
            --go_out=. \
            --go_opt=paths=source_relative \
            "$proto"
    done
    
    # Generate TypeScript files
    if [ ! -d "node_modules" ]; then
        echo "  Installing npm dependencies..."
        npm i
    else
        echo "  npm dependencies already installed"
    fi
    
    for proto in $(find . -name "*.proto" -not -path "../node_modules/*" -not -path "./build/*"); do
        echo "  Generating TypeScript files for: $proto"
        protoc \
            --plugin=protoc-gen-ts_proto="../node_modules/.bin/protoc-gen-ts_proto" \
            --ts_proto_out=. \
            --ts_proto_opt=esModuleInterop=true \
            --proto_path=. \
            --proto_path="$(go env GOPATH)/src" \
            "$proto"
    done
    
    # Build TypeScript
    echo "  Building TypeScript..."
    npm run build
    
    # Git add
    git add build/
    # Add all TypeScript files (including in subdirectories, excluding build and node_modules)
    find . -name "*.ts" -not -path "./build/*" -not -path "./node_modules/*" -exec git add {} \;
    
    echo "✓ Build complete for all models"
}

# Main menu
echo "=========================================="
echo "  Model Build Script"
echo "=========================================="
echo ""
echo "Select build option:"
echo "  1) Build all models"
echo "  2) Build a specific model"
echo ""
read -p "Enter your choice (1 or 2): " choice

case $choice in
    1)
        build_all
        ;;
    2)
        # Get model directories
        model_dirs=($(get_model_dirs))
        
        if [ ${#model_dirs[@]} -eq 0 ]; then
            echo "Error: No model directories found"
            exit 1
        fi
        
        echo ""
        echo "Available models:"
        echo ""
        
        # Display menu
        index=1
        valid_models=()
        for dir in "${model_dirs[@]}"; do
            if has_proto_files "$dir"; then
                echo "  $index) $dir"
                valid_models+=("$dir")
                ((index++))
            fi
        done
        
        if [ ${#valid_models[@]} -eq 0 ]; then
            echo "Error: No model directories with .proto files found"
            exit 1
        fi
        
        echo ""
        read -p "Select a model (1-${#valid_models[@]}): " model_choice
        
        # Validate input
        if ! [[ "$model_choice" =~ ^[0-9]+$ ]] || [ "$model_choice" -lt 1 ] || [ "$model_choice" -gt ${#valid_models[@]} ]; then
            echo "Error: Invalid selection"
            exit 1
        fi
        
        selected_model="${valid_models[$((model_choice - 1))]}"
        echo ""
        build_model "$selected_model"
        ;;
    *)
        echo "Error: Invalid choice. Please select 1 or 2"
        exit 1
        ;;
esac

# Cleanup
if [ -d "../node_modules" ]; then
    rm -rf ../node_modules
fi

echo ""
echo "Done!"
