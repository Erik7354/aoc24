const std = @import("std");

const print = @import("std").debug.print;
const panic = @import("std").debug.panic;

const day1_file = "day1.txt";
const nLines = 1000;

pub fn day1() u32 {
    const file = std.fs.cwd().openFile(day1_file, .{}) catch unreachable;
    defer file.close();

    // Create a buffered reader for the file
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

    std.mem.sort(u32, arr1, {}, comptime std.sort.asc(u32));
    std.mem.sort(u32, arr2, {}, comptime std.sort.asc(u32));

    var distance: u32 = 0;
    for (0..arr1.len) |ii| {
        distance += @max(arr1[ii], arr2[ii]) - @min(arr1[ii], arr2[ii]);
    }

    print("Distance: {any}\n", .{distance});

    return distance;
}

// zig test --test-filter "day1 test" day1.zig
test "day1 test" {
    try std.testing.expect(day1() == 1660292);
}
