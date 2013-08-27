package model;

/**
 * Represents a cell on the grid. Can be used to convey where the 
 * user selected, or as part of a specification of a board.
 */
public class Cell {

    // The character at the specified cell.
    protected char m_letter;

    // The row and column indices of the cell on its board.
    protected int m_row;
    protected int m_column;

    public Cell(int row, int column, char letter)
    {
        if (! Character.isLetter(letter)) {
            throw new IllegalArgumentException("The given character is not a letter.");
        }

        if (row < 0 || column < 0)
            throw new IllegalArgumentException("Board indices may not be negative.");

        m_letter = letter;
        m_row = row;
        m_column = column;
    }

    public int getRow()
    {
        return m_row;
    }

    public int getColumn()
    {
        return m_column;
    }

    public char getLetter()
    {
        return m_letter;
    }

    /**
     * Determines if the given cell is contiguous to this cell.
     */
    public boolean isContiguous(Cell that)
    {
        return Math.abs(that.getRow() - m_row) <= 1 &&
               Math.abs(that.getColumn() - m_column) <= 1;
    }

    @Override
    public boolean equals(Object other)
    {
        if (other == null || !(other instanceof Cell))
            return false;

        Cell otherCell = (Cell) other;

        return otherCell.getRow() == m_row &&
               otherCell.getColumn() == m_column &&
               otherCell.getLetter() == m_letter;
    }

    @Override
    public int hashCode()
    {
        int result = 14;
        result = result * 37 + m_row;
        result = result * 37 + m_column;
        result = result * 37 + (int) m_letter;

        return result;
    }

}
