local numbers_ascii = {
    [0] = {
        "  #####  ",
        " #     # ",
        " #     # ",
        " #     # ",
        " #     # ",
        " #     # ",
        "  #####  "
    },
    [1] = {
        "    #    ",
        "   ##    ",
        "  # #    ",
        "    #    ",
        "    #    ",
        "    #    ",
        " ####### "
    },
    [2] = {
        " ######  ",
        "      #  ",
        "      #  ",
        " ######  ",
        " #       ",
        " #       ",
        " ####### "
    },
    [3] = {
        " ######  ",
        "      #  ",
        "      #  ",
        " ######  ",
        "      #  ",
        "      #  ",
        " ######  "
    },
    [4] = {
        " #     # ",
        " #     # ",
        " #     # ",
        " ####### ",
        "      #  ",
        "      #  ",
        "      #  "
    },
    [5] = {
        " ####### ",
        " #       ",
        " #       ",
        " ######  ",
        "      #  ",
        "      #  ",
        " ######  "
    },
    [6] = {
        "  #####  ",
        " #       ",
        " #       ",
        " ######  ",
        " #     # ",
        " #     # ",
        "  #####  "
    },
    [7] = {
        " ####### ",
        "      #  ",
        "     #   ",
        "    #    ",
        "   #     ",
        "  #      ",
        "  #      "
    },
    [8] = {
        "  #####  ",
        " #     # ",
        " #     # ",
        "  #####  ",
        " #     # ",
        " #     # ",
        "  #####  "
    },
    [9] = {
        "  #####  ",
        " #     # ",
        " #     # ",
        "  ###### ",
        "      #  ",
        "      #  ",
        "  #####  "
    }
}

local function draw_number(buffer, dim, num)
    local height = dim.height
    local width = dim.width
    local num_str = tostring(num)
    local max_digit_width = 5 -- Assuming each digit has width of 5
    local max_digit_height = 7 -- Maximum height of a single digit
    local x = math.floor((width - #num_str * max_digit_width) / 2)
    local y = math.floor((height - max_digit_height) / 2)

    for i = 1, #num_str do
        local digit = tonumber(num_str:sub(i, i))
        local ascii_art = numbers_ascii[digit]

        if ascii_art then
            local digit_width = #ascii_art[1]
            local digit_height = #ascii_art

            -- Ensure the width of the ASCII art does not exceed the remaining width in the buffer
            if x + digit_width > width then
                -- Adjust x to fit the remaining space
                x = width - digit_width + 1
            end

            -- Ensure the height of the ASCII art does not exceed the remaining height in the buffer
            if y + digit_height > height then
                -- Adjust y to fit the remaining space
                y = height - digit_height + 1
            end

            -- Draw the digit onto the buffer
            for k = 1, digit_height do
                for j = 1, digit_width do
                    local char = ascii_art[k]:sub(j, j)
                    if char ~= ' ' then
                        buffer:PlacePoint(x + j - 1, y + k - 1, char)
                    end
                end
            end

            x = x + digit_width + 1
        else
            print("Number not supported")
            return
        end
    end
end

return {draw_score = draw_number}
