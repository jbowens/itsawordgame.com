(function() {

    // Setup
    $(document).ready(function(e) {

        constructRandomBoard();

        // TODO: Do setup
    });

    var width = 6;
    var height = 5;

    function constructCell(letter, row, column) {
        var cell = $('<div class="cell"></div>');
        cell.data('row', row);
        cell.data('column', column);

        var cellOverlay = $('<div class="overlay"></div>');
        cell.append(cellOverlay);

        var innerCell = $('<div class="inner-cell"></div>');
        innerCell.text(letter);
        cell.append(innerCell);
        cell.hover(cellHover, cellExit);
        return cell;
    }

    function cellHover(e) {
        $(this).addClass('hover');
    }

    function cellExit(e) {
        var cell = $(this);
        var overlay = cell.find('.overlay');
        cell.removeClass('hover');
        // Do an exit fade animation

        if (cell.data('animationTimeout')) {
            clearTimeout(cell.data('animationTimeout'));
            cell.data('animationTimeout', null);
        }

        var num = 20;
        function animationTick() {
            var alpha = num / 20.0;
            var color = 'rgba(163,67,99,'+alpha+')';
            console.log(overlay);
            overlay.css('background-color', color);
            num--;

            if (num > 0)
                cell.data('animationTimeout', setTimeout(animationTick, 50));
        }

        animationTick();

    }

    function constructRandomBoard() {
        var board = $("#board");

        for (var r = 0; r < height; r++)
        {
            for (var c = 0; c < width; c++)
            {
                var randomNumber = Math.floor(Math.random() * 26);
                var randomLetter = String.fromCharCode(randomNumber + 65);
                var cell = constructCell(randomLetter, r, c);
                board.append(cell);
            }
        }
    }

})();
