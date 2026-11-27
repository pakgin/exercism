const std = @import("std");
const mem = std.mem;

pub fn factors(allocator: mem.Allocator, value: u64) mem.Allocator.Error![]u64 {
    var ans: std.ArrayList(u64) = .empty;
    defer ans.deinit(allocator);

    var n = value;
    var d: u64 = 2;
    // Only need to try divisors up to sqrt(n); `d <= n / d` is `d * d <= n`
    // without the risk of overflowing.
    while (d <= n / d) {
        if (n % d == 0) {
            try ans.append(allocator, d);
            n /= d; // keep the same d: it may divide n again (e.g. 8 = 2*2*2)
        } else {
            d += 1;
        }
    }
    // Whatever is left over (if > 1) is itself a prime factor.
    if (n > 1) try ans.append(allocator, n);

    return ans.toOwnedSlice(allocator);
}
