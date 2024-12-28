const std = @import("std");

const print = @import("std").debug.print;
const panic = @import("std").debug.panic;

const day1_plus_file = "day1.txt";
const nLines = 1000;

fn day1_plus() u64 {
    const file = std.fs.cwd().openFile(day1_plus_file, .{}) catch unreachable;
    defer file.close();

    var reader = std.io.bufferedReader(file.reader());
    const r = reader.reader();

    var allocator = std.heap.GeneralPurposeAllocator(.{}){};
    const allo = allocator.allocator();

    const chunk = allo.alloc(u8, 15) catch unreachable; // memory is not initialized
    const arr1 = allo.alloc(u32, nLines) catch unreachable;
    const arr2 = allo.alloc(u32, nLines) catch unreachable;

    var i: u32 = 0;
    while (r.readUntilDelimiterOrEof(chunk[0..], '\n') catch unreachable) |line| {
        var split = std.mem.split(u8, line, "   ");
        const int1 = std.fmt.parseInt(u32, split.first(), 10) catch unreachable;
        const int2 = std.fmt.parseInt(u32, split.next().?, 10) catch unreachable;

        arr1[i] = int1;
        arr2[i] = int2;

        i += 1;
    }

    var score: u64 = 0;
    for (0..arr1.len) |ii| {
        for (0..arr2.len) |jj| {
            if (arr1[ii] == arr2[jj]) {
                score += arr1[ii];
            }
        }
    }

    print("Score: {any}\n", .{score});
}

// zig test --test-filter "day1+ test" day1+.zig
test "day1+ test" {
    try std.testing.expect(day1_plus() == 22776016);
}
