#!/usr/bin/env bash

echo "==> Checking that code complies with gocritic requirements..."
gocritic check-project --enable=all -withExperimental -withOpinionated .
if [ $? -eq 1 ]; then
    echo ""
    echo "Gocritic found suspicious constructs. Please check the reported constructs"; \
    echo "and fix them if necessary before submitting the code for review."; \
    exit 1
fi

exit 0