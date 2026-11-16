const std = @import("std");

pub fn isArmstrongNumber(num: u128) bool {
    var curr = num;
    const nDigit, var sum: u128 = .{ countDigits(curr), 0 };
    while (curr > 0) {
        const digit = curr % 10;
        sum += std.math.pow(u128, digit, nDigit);
        curr /= 10;
    }
    return sum == num;
}

fn countDigits(num: u128) u8 {
    if (num == 0) return 1;
    return std.math.log10_int(num) + 1;
}
