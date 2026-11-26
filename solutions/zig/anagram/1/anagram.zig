const std = @import("std");
const mem = std.mem;

/// Returns the set of strings in `candidates` that are anagrams of `word`.
/// Caller owns the returned memory.
pub fn detectAnagrams(
    allocator: mem.Allocator,
    word: []const u8,
    candidates: []const []const u8,
) !std.BufSet {
    var ans = std.BufSet.init(allocator);
    errdefer ans.deinit();

    const wordSorted = try sortedLower(allocator, word);
    defer allocator.free(wordSorted);

    for (candidates) |c| {
        if (c.len != word.len or std.ascii.eqlIgnoreCase(word, c)) continue;

        const cSorted = try sortedLower(allocator, c);
        defer allocator.free(cSorted);

        if (mem.eql(u8, wordSorted, cSorted)) {
            try ans.insert(c);
        }
    }
    return ans;
}

fn sortedLower(allocator: mem.Allocator, s: []const u8) ![]u8 {
    const buf = try allocator.alloc(u8, s.len);
    _ = std.ascii.lowerString(buf, s);
    mem.sort(u8, buf, {}, std.sort.asc(u8));
    return buf;
}
