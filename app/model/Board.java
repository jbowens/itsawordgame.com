package model;

import java.util.Collection;

/**
 * Represents an arrangement of letters on a board.
 */
public class Board
{
    // The dimensions of the board
    protected int m_width;
    protected int m_height;

    // List of cell information
    protected Collection<Cell> m_cells;

    public Board(int width, int height, Collection<Cell> cells)
    {
        m_width = width;
        m_height = height;
        m_cells = cells;
    }


}
