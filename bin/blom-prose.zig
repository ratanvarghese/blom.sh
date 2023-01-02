const std = @import("std");

const JsfItem = struct { url: []u8, title: []u8, tags: [][]u8 };

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

    for (items) |item| {
        for (item.tags) |tag| {
            if (std.mem.eql(u8, tag, "joy of writing")) {
                try stdout.print("\t<a class=\"creative-box creative-prose\" href=\"{s}\">\n", .{item.url});
                try stdout.print("\t\t<img src=\"{s}/attachments/thumbnail.svg\" />\n", .{item.url});
                try stdout.print("\t\t<label>{s}</label>\n", .{item.title});
                try stdout.print("\t</a>\n", .{});
                break;
            }
        }
    }
}
