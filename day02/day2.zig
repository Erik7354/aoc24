const std = @import("std");

const print = @import("std").debug.print;
const panic = @import("std").debug.panic;

const filepath = "day2.txt";

const minGrowth = 1;
const maxGrowth = 3;

fn day2() u16 {
    const file = std.fs.cwd().openFile(filepath, .{}) catch unreachable;
    defer file.close();

    var reader = std.io.bufferedReader(file.reader());
    const r = reader.reader();

    var gpallocator = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpallocator.allocator();

    const chunk = allocator.alloc(u8, 100) catch unreachable;
    defer allocator.free(chunk);

    var nSafeReports: u16 = 0;
    while (r.readUntilDelimiterOrEof(chunk, '\n') catch unreachable) |line| {
        print("{s}\n", .{line});
        var si = std.mem.split(u8, line, " ");

        // parse nums
        var nums = std.ArrayList(isize).init(allocator);
        while (si.next()) |snum| {
            nums.append(std.fmt.parseInt(isize, snum, 10) catch unreachable) catch unreachable;
        }

        if (checkNums(nums)) {
            nSafeReports += 1;
        }
    }

    print("safe reports: {d}\n", .{nSafeReports});
    return nSafeReports;
}

fn checkNums(al: std.ArrayList(isize)) bool {
    var ascending: ?bool = null;

    for (0..al.items.len - 1) |i| {
        const lastNum = al.items[i];
        const num = al.items[i + 1];
        const diff = lastNum - num;

        if (ascending == null) {
            ascending = diff > 0;
        }

        if (ascending.? and lastNum < num) {
            return false;
        }
        if (!ascending.? and lastNum > num) {
            return false;
        }
        if (@abs(lastNum - num) < minGrowth) {
            return false;
        }
        if (@abs(lastNum - num) > maxGrowth) {
            return false;
        }
    }

    return true;
}

// zig test --test-filter "day2 test" day2.zig
test "day2 test" {
    try std.testing.expect(day2() == 660);
}
