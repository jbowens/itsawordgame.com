package model;

import java.util.Collection;

/**
 * Represents a round of the game. This contains both the board and the associated
 * word data.
 *
 * @author jbowens
 */
public class Round
{
    /* The board arrangement for this round. */
    private Board m_board;

    public Round(Board board, Collection<Word> words)
    {
        // TODO: Use words
        m_board = board;
    }

    /**
     * Retrieve's the round's board.
     */
    public Board getBoard()
    {
        return m_board;
    }

    /**
     * Creates a word discoverer for finding words in letter
     * sequences.
     */
    public WordDiscoverer createWordDiscoverer()
    {
        // TODO: Implement
        return null;
    }

}
