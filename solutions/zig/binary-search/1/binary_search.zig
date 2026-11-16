// Take a look at the tests, you might have to change the function arguments

pub fn binarySearch(comptime T: type, target: T, items: []const T) ?usize {
    var left: usize = 0;
    var right: usize = items.len;
    while (left < right) {
        const mid = (left + right) / 2;
        if (items[mid] == target) {
            return mid;
        } else if (items[mid] < target) {
            left = mid + 1;
        } else {
            right = mid;
        }
    }
    return null;
}
