#!/bin/bash

# Setup script for git hooks
# This script installs the pre-commit hooks for the project

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}  Git Hooks Setup${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo ""

# Get the git root directory
GIT_ROOT=$(git rev-parse --show-toplevel 2>/dev/null)

if [ -z "$GIT_ROOT" ]; then
    echo -e "${RED}❌ Error: Not in a git repository${NC}"
    exit 1
fi

cd "$GIT_ROOT"

# Configure git to use our hooks directory
echo "📁 Setting up hooks directory..."
git config core.hooksPath .githooks

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Git hooks directory configured${NC}"
else
    echo -e "${RED}❌ Failed to configure git hooks directory${NC}"
    exit 1
fi

# Make sure all hooks are executable
echo "🔧 Making hooks executable..."
if [ -d ".githooks" ]; then
    chmod +x .githooks/* 2>/dev/null || true
    echo -e "${GREEN}✅ Hooks are executable${NC}"
else
    echo -e "${YELLOW}⚠️  No .githooks directory found${NC}"
fi

echo ""
echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${GREEN}✅ Git hooks setup complete!${NC}"
echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo ""
echo "The following hooks are now active:"
echo "  • pre-commit: Runs tests, linting, and formatting checks"
echo ""
echo "To run checks manually before committing:"
echo -e "  ${YELLOW}make check-all${NC}     - Run all checks"
echo -e "  ${YELLOW}make check-quick${NC}  - Run only linting and formatting"
echo -e "  ${YELLOW}make fix-all${NC}      - Auto-fix issues"
echo ""
echo "To skip pre-commit checks (not recommended):"
echo -e "  ${YELLOW}SKIP_CHECKS=1 git commit ...${NC}"
echo ""
echo "To disable hooks for this repository:"
echo -e "  ${YELLOW}git config --unset core.hooksPath${NC}"