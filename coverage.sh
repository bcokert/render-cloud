#!/usr/bin/env bash

COVERAGE_DIR=.coverage-reports
profile="$COVERAGE_DIR/cover.out"
mode=count

generate_cover_data() {
    rm -rf "$COVERAGE_DIR"
    mkdir "$COVERAGE_DIR"

    for pkg in "$@"; do
        f="$COVERAGE_DIR/$(echo $pkg | tr / -).cover"
        go test -covermode="$mode" -coverprofile="$f" "$pkg"
    done

    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$COVERAGE_DIR"/*.cover >>"$profile"
}

show_cover_report() {
    go tool cover -html="$profile"
}

generate_cover_data $(go list ./...)
show_cover_report
