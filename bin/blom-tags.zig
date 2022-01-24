const std = @import("std");

const JsfItem = struct { url: []u8, title: []u8, tags: [][]u8 };

fn sliceLessThan(comptime T: type) fn ([][]T, []T, []T) bool {
    const impl = struct {
        fn inner(_: [][]T, a: []T, b: []T) bool {
            for (b) |char_b, i| {
                var char_a: u8 = 'z';
                if (i < a.len) {
                    char_a = a[i];
                }
                if (char_a < char_b) {
                    return true;
                }
                if (char_a > char_b) {
                    return false;
                }
            }
            return false;
        }
    };

    return impl.inner;
}

fn title(s: []u8, allocator: std.mem.Allocator) ![]u8 {
    var res = try allocator.alloc(u8, s.len);
    const diff: u8 = 'a' - 'A';
    for (s) |_, i| {
        if (i == 0 or s[i - 1] == ' ') {
            res[i] = s[i] - diff;
        } else {
            res[i] = s[i];
        }
    }
    return res;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    const stdin = std.io.getStdIn().reader();
    const stdout = std.io.getStdOut().writer();

    const input = try stdin.readAllAlloc(allocator, 1000000);
    var stream = std.json.TokenStream.init(input[0..]);
    const items = try std.json.parse([]JsfItem, &stream, .{
        .allocator = allocator,
        .ignore_unknown_fields = true,
    });

    var tag_map = std.StringHashMap([]bool).init(allocator);
    var tag_list = std.ArrayList([]u8).init(allocator);

    for (items) |item, i| {
        for (item.tags) |tag| {
            if (!tag_map.contains(tag)) {
                var new_list = try allocator.alloc(bool, items.len);
                for (new_list) |*b| {
                    b.* = false;
                }
                try tag_map.put(tag, new_list);
                try tag_list.append(tag);
            }
            var list = tag_map.get(tag) orelse unreachable;
            list[i] = true;
        }
    }

    std.sort.sort([]u8, tag_list.items, tag_list.items, comptime sliceLessThan(u8));

    for (tag_list.items) |tag| {
        const title_tag = title(tag, allocator);
        try stdout.print("### {s}\n", .{title_tag});
        const list = tag_map.get(tag) orelse unreachable;
        for (list) |item, item_i| {
            if (item) {
                try stdout.print("+ [{s}]({s})\n", .{
                    items[item_i].title, items[item_i].url,
                });
            }
        }
        try stdout.print("\n", .{});
    }
}
