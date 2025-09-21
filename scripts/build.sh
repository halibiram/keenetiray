#!/bin/bash

# A script to build and package the application for different target architectures.
# Usage: ./scripts/build.sh <architecture>
# Example: ./scripts/build.sh arm

set -e

# --- Configuration ---
TARGET_ARCH=$1
if [ -z "$TARGET_ARCH" ]; then
    echo "Usage: $0 <architecture>"
    echo "Supported architectures: arm, mips, amd64 (for local testing)"
    exit 1
fi

TARGET_OS="linux"
DIST_DIR="dist/${TARGET_OS}_${TARGET_ARCH}"
GO_BINARY_PATH="$DIST_DIR/v2ray-keenetic"
ARCHIVE_NAME="v2ray-keenetic_${TARGET_OS}_${TARGET_ARCH}.tar.gz"

echo "--- Starting V2Ray Keenetic Build Process for ${TARGET_OS}/${TARGET_ARCH} ---"

# 1. Build Frontend (if necessary)
if [ ! -d "web/dist" ]; then
    echo ">>> Frontend build artifacts not found. Building Frontend..."
    (cd web && npm install && npm run build)
    echo ">>> Frontend build complete."
else
    echo ">>> Frontend build artifacts found. Skipping frontend build."
fi

# 2. Build Backend for Target Architecture
echo ">>> Building Backend for ${TARGET_OS}/${TARGET_ARCH}..."
rm -rf $DIST_DIR # Clean up old build for this specific target
mkdir -p $DIST_DIR

echo "Building binary at $GO_BINARY_PATH"
CGO_ENABLED=0 GOOS=$TARGET_OS GOARCH=$TARGET_ARCH go build -ldflags="-s -w" -o $GO_BINARY_PATH ./cmd/v2ray-keenetic
echo ">>> Backend build complete."

# 3. Stage Files
echo ">>> Staging files for distribution..."
cp -r web/dist $DIST_DIR/web

# 4. Create final .tar.gz archive
echo ">>> Creating distributable archive..."
RELEASE_DIR="releases"
mkdir -p $RELEASE_DIR
(
  cd $DIST_DIR && \
  tar -czf ../../$RELEASE_DIR/$ARCHIVE_NAME .
)
echo ">>> Archive created at $RELEASE_DIR/$ARCHIVE_NAME"

echo "--- Build and packaging process complete. ---"
echo "Final archive is available at $RELEASE_DIR/$ARCHIVE_NAME"
