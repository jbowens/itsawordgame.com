(function() {

    // Setup
    $(document).ready(function(e) {

        constructRandomBoard();

        // TODO: Do setup
    });

    var width = 5;
    var height = 4;

    function constructCell(letter, row, column) {
        var cell = $('<div class="cell"></div>');
        cell.data('row', row);
        cell.data('column', column);

        var innerCell = $('<div class="inner-cell"></div>');
        var letterSpan = $('<span class="letter"></span>');
        letterSpan.text(letter);
        innerCell.append(letterSpan);
        cell.append(innerCell);
        letterSpan.hover(cellHover, cellExit);
        return cell;
    }

    function cellHover(e) {
        var cell = $(this).closest('.cell');
        cell.addClass('hover');
    }

    function cellExit(e) {
        var cell = $(this).closest('.cell');
        var innerCell = cell.find('.inner-cell');
        cell.removeClass('hover');

        // Do an exit fade animation
        if (cell.data('animationTimeout')) {
            clearTimeout(cell.data('animationTimeout'));
            cell.data('animationTimeout', null);
        }

        var num = 30;
        function animationTick() {
            var alpha = num / 30.0;
            var color = 'rgba(163,67,99,'+alpha+')';
            innerCell.css('background-color', color);
            num--;

            if (num > 0)
                cell.data('animationTimeout', setTimeout(animationTick, 30));
        }

        animationTick();

    }

    function constructRandomBoard() {
        var board = $("#board");
        var clearfix = $("#board .clearfix");

        for (var r = 0; r < height; r++)
        {
            for (var c = 0; c < width; c++)
            {
                var randomNumber = Math.floor(Math.random() * 26);
                var randomLetter = String.fromCharCode(randomNumber + 65);
                var cell = constructCell(randomLetter, r, c);
               $(cell).insertBefore(clearfix); 
            }
        }
    }

})();
