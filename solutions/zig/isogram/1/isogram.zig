const std = @import("std");

pub fn isIsogram(str: []const u8) bool {
    var seen = std.AutoHashMap(u8, void).init(std.heap.page_allocator);
    defer seen.deinit();

    for (str) |c| {
        if (!std.ascii.isAlphabetic(c)) continue;
        const gop = seen.getOrPut(std.ascii.toLower(c)) catch return false;
        if (gop.found_existing) return false;
    }
    return true;
}
