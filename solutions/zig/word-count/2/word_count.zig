const std = @import("std");
const mem = std.mem;

/// Every byte that separates words: anything except a letter, digit or apostrophe.
const separators = blk: {
    var buf: [256]u8 = undefined;
    var len: usize = 0;
    for (0..256) |i| {
        const c: u8 = @intCast(i);
        if (!std.ascii.isAlphanumeric(c) and c != '\'') {
            buf[len] = c;
            len += 1;
        }
    }
    const result = buf[0..len].*;
    break :blk &result;
};

/// Returns the counts of the words in `s`.
/// Caller owns the returned map and its keys.
pub fn countWords(allocator: mem.Allocator, s: []const u8) !std.StringHashMap(u32) {
    var map = std.StringHashMap(u32).init(allocator);
    errdefer freeMap(allocator, &map);

    var tokens = mem.tokenizeAny(u8, s, separators);
    while (tokens.next()) |token| {
        // Apostrophes at the edges are quotation marks, not part of the word.
        const trimmed = mem.trim(u8, token, "'");
        if (trimmed.len == 0) continue;

        const word = try std.ascii.allocLowerString(allocator, trimmed);
        errdefer allocator.free(word);

        const entry = try map.getOrPut(word);
        if (entry.found_existing) {
            allocator.free(word);
            entry.value_ptr.* += 1;
        } else {
            entry.value_ptr.* = 1;
        }
    }

    return map;
}

fn freeMap(allocator: mem.Allocator, map: *std.StringHashMap(u32)) void {
    var keys = map.keyIterator();
    while (keys.next()) |key| allocator.free(key.*);
    map.deinit();
}
