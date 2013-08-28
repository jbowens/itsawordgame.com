(function() {

    // Constants
    var ANIMATION_GRANULARITY = 30;
    var ANIMATION_PAUSE_MILLI = 40;

    // Setup
    $(document).ready(function(e) {

        initGame();

        // TODO: Do setup
    });

    var width = 5;
    var height = 4;

    function initGame() {
        $.get('/init-game', function(round) {
            round = $.parseJSON(round);
            var cells = round.board.cells;

            // Sort the cells by row, column
            cells.sort(function (a, b) { return a.row * round.board.width + a.column - (b.row * round.board.width + b.column)});

            var clearfix = $("#board .clearfix");

            // Empty out the board
            clearfix.siblings().remove();
    
            // Add the new cells
            for (var i = 0; i < cells.length; i++)
            {
                var cell = constructCell(cells[i].letter, cells[i].row, cells[i].column);
                $(cell).insertBefore(clearfix); 
            }
        });
    }

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

        // Fade from the hover color to the regular cell background
        fade(innerCell, {r: 163, g:67, b:99});
    }

    /* Animates a color change for the given element, from the given color to
     * transparency.
     */
    function fade(el, color)
    {
        el = $(el);
        // If there's an existing animation, cancel it. This newer
        // one takes precedence.
        if (el.data('animationTimeout')) {
            clearTimeout(el.data('animationTimeout'));
            el.data('animationTimeout', null);
        }
        
        var animationFrame = ANIMATION_GRANULARITY;
        function animationTick() {
            var alpha = animationFrame / ANIMATION_GRANULARITY;
            var cssColor = 'rgba('+color.r+','+color.g+','+color.b+','+alpha+')';
            el.css('background-color', cssColor);
            animationFrame--;

            if (animationFrame > 0)
                el.data('animationTimeout', setTimeout(animationTick, ANIMATION_PAUSE_MILLI));
        }
        animationTick();
    }

})();
