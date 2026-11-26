const std = @import("std");
const mem = std.mem;

fn isWordByte(c: u8) bool {
    return std.ascii.isAlphanumeric(c);
}

/// True only for an apostrophe with a letter or digit on both sides.
fn isInnerApostrophe(s: []const u8, i: usize) bool {
    if (s[i] != '\'') return false;
    if (i == 0 or i + 1 >= s.len) return false;
    return isWordByte(s[i - 1]) and isWordByte(s[i + 1]);
}

/// Returns the counts of the words in `s`.
/// Caller owns the returned memory.
pub fn countWords(allocator: mem.Allocator, s: []const u8) !std.StringHashMap(u32) {
    var map = std.StringHashMap(u32).init(allocator);
    errdefer {
        var iter = map.keyIterator();
        while (iter.next()) |key_ptr| {
            allocator.free(key_ptr.*);
        }
        map.deinit();
    }

    var i: usize = 0;
    while (i < s.len) {
        while (i < s.len and !isWordByte(s[i])) : (i += 1) {}
        if (i == s.len) break;

        const start = i;
        while (i < s.len and (isWordByte(s[i]) or isInnerApostrophe(s, i))) : (i += 1) {}

        const owned = try allocator.dupe(u8, s[start..i]);
        for (owned) |*c| c.* = std.ascii.toLower(c.*);

        const gop = map.getOrPut(owned) catch |err| {
            allocator.free(owned);
            return err;
        };
        if (gop.found_existing) {
            allocator.free(owned);
            gop.value_ptr.* += 1;
        } else {
            gop.value_ptr.* = 1;
        }
    }

    return map;
}
