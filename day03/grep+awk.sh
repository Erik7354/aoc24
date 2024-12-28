grep -o -E "mul\([0-9]+,[0-9]+\)" in.txt | awk 'BEGIN { FS=","; sum = 0 } { gsub(/mul\(/, "", $0); gsub(/\)/, "", $0); sum = sum + $1*$2 } END { print sum }'

