import os
import sys

UNITS = {
    "B": 1,
    "KB": 1024,
    "MB": 1024 ** 2,
    "GB": 1024 ** 3,
}


def parse_size(size_str):
    size_str = size_str.strip().upper()

    for unit in ["GB", "MB", "KB", "B"]:
        if size_str.endswith(unit):
            number = float(size_str[:-len(unit)])
            return int(number * UNITS[unit])

    # Assume raw bytes if no unit provided
    return int(size_str)


def generate_random_file(filename, size_bytes):
    chunk_size = 1024 * 1024  # 1 MB chunks

    with open(filename, "wb") as f:
        remaining = size_bytes

        while remaining > 0:
            current_chunk = min(chunk_size, remaining)
            f.write(os.urandom(current_chunk))
            remaining -= current_chunk

    print(f"Generated '{filename}' with size {size_bytes} bytes")


def main():
    if len(sys.argv) != 3:
        print("Usage: python generate_random_file.py <filename> <size>")
        print("Examples:")
        print("  python generate_random_file.py file.bin 1048576")
        print("  python generate_random_file.py file.bin 10MB")
        sys.exit(1)

    filename = sys.argv[1]
    size_bytes = parse_size(sys.argv[2])

    generate_random_file(filename, size_bytes)


if __name__ == "__main__":
    main()