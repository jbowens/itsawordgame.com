package model;

/**
 * Generates new rounds... 
 *
 * @author jbowens
 */
public class RoundGenerator
{
    public static final BoardFactory DEFAULT_BOARD_FACTORY = new DistributionBoardFactory();

    // The board dimensions with which to construct boards.
    protected final int m_boardWidth;
    protected final int m_boardHeight;

    // The board factory to use for producing new boards.
    protected final BoardFactory m_boardFactory;

    public RoundGenerator(int boardWidth, int boardHeight)
    {
        m_boardWidth = boardWidth;
        m_boardHeight = boardHeight;
        m_boardFactory = DEFAULT_BOARD_FACTORY;
    }

    public Round generateRandomRound()
    {
        // Generate the board first. 
        Board board = m_boardFactory.generateBoard(m_boardWidth, m_boardHeight);

        Round round = new Round(board, new java.util.ArrayList<Word>());
        return round;
    }

}
