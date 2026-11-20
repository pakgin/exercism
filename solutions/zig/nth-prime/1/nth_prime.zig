const std = @import("std");
const mem = std.mem;

pub fn prime(allocator: mem.Allocator, number: usize) !usize {
    _ = allocator;
    if (number == 0) return error.ZeroNotAllowed;
    var i: usize = 0;
    var candidate: usize = 1;
    while (i < number) {
        candidate += 1;
        if (isPrime(candidate)) i += 1;
    }
    return candidate;
}

fn isPrime(number: usize) bool {
    var i: usize = 2;
    while (i * i <= number) : (i += 1) {
        if (number % i == 0) return false;
    }
    return true;
}
