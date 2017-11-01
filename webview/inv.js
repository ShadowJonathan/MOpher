// [@POS [X,Y]]
pos = [
    [462, 84], // CRAFT END

    [294, 54], // CRAFT
    [348, 54],
    [294, 108],
    [348, 108],

    [24, 24], // ARMOR
    [24, 78],
    [24, 132],
    [24, 186],

    [24, 252], // INVENTORY, FIRST ROW
    [78, 252],
    [132, 252],
    [186, 252],
    [240, 252],
    [294, 252],
    [348, 252],
    [402, 252],
    [456, 252],
    [24, 306], // SECOND ROW
    [78, 306],
    [132, 306],
    [186, 306],
    [240, 306],
    [294, 306],
    [348, 306],
    [402, 306],
    [456, 306],
    [24, 360], // THIRD ROW
    [78, 360],
    [132, 360],
    [186, 360],
    [240, 360],
    [294, 360],
    [348, 360],
    [402, 360],
    [456, 360],

    [24, 426], // HOTBAR
    [78, 426],
    [132, 426],
    [186, 426],
    [240, 426],
    [294, 426],
    [348, 426],
    [402, 426],
    [456, 426],

    [231, 186] // SECOND HAND
];

cursorPos = [126, 153];

// type: X, Y
typeIndex = {
    chest: [7, 39],
    cobblestone: [10, 27],
    dirt: [13, 30],
    sand: [3, 31],
    mossy_cobblestone: [25, 30],
    gold_ore: [24, 31],
    gravel: [21, 30],
    red_sand: [1, 31],
};

$(() => {
    let V = $("#view");
    for (let i in pos) {
        let p = pos[i];
        V.append(`<div class="item" style="top: ${p[1]}px; left: ${p[0]}px;" id="${i}"><div class='tl'>${i}</div><div class='br hidden'></div></div>`)
        $(`#${i}`).on("click", () => {
            sock.send(i + "")
        })
    }
    V.append(`<div class="item" style="top: ${cursorPos[1]}px; left: ${cursorPos[0]}px;" id="-1"><div class='tl'>-1</div><div class='br hidden'></div></div>`)
    $(`#-1`).on("click", () => {
        sock.send("-1")
    })
});

/**
 *
 * @param {Array.<{Amount: number, Type: string}|null>} entries
 * @param {{Amount: number, Type: string}|null} onCursor
 */
function process(entries, onCursor) {
    let V = $(".item");
    V.each((_, e) => {
        e = $(e);
        let i = parseInt(e.attr('id'));
        let br = e.find(".br");
        if (i >= 0) {
            if (entries[i]) {
                let v = entries[i];
                doImage(v.Type, e);
                br.html(v.Amount);
                br.toggleClass("hidden", false);
            } else {
                e.css("background-image", "");
                e.css("background-position", "");
                br.toggleClass("hidden", true)
                e.toggleClass("all-icon", false)
            }
        } else {
            if (onCursor) {
                let v = onCursor;
                doImage(v.Type, e);
                br.toggleClass("hidden", false);
                br.html(v.Amount)
            } else {
                e.css("background-image", "");
                e.css("background-position", "");
                br.toggleClass("hidden", true);
                e.toggleClass("all-icon", false)
            }
        }
    })
}

function doImage(type, e) {
    if (!typeIndex[type]) {
        e.css("background-image", `url(img/views/${type}.png`);
        e.css("background-position", "");
        e.toggleClass("all-icon", false)
    } else {
        e.css("background-image", `url(img/all.png`);
        e.css("background-position", `-${typeIndex[type][0] * 48}px -${typeIndex[type][1] * 48}px`);
        e.toggleClass("all-icon", true)
    }
}