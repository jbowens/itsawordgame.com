package model

import java.util.ArrayList;
import java.util.Collection;
import java.util.Random;

/**
 * Constucts boards using the letter frequency of the English language.
 *
 * @author jbowens
 */
public class DistributionBoardFactory
{

    /* Frequencies for all 26 letters of the alphabet.
     */
    protected static final double[] LETTER_FREQUENCIES = {.08167, .01492, .02782, .04253, .12702, .02228,
                                                          .02015, .06094, .06966, .00153, .00772, .04025,
                                                          .02406, .06749, .07507, .01929, .00095, .05987,
                                                          .06327, .09056, .02758, .00978, .02360, .00150,
                                                          .01974, .00074}

    // Random number generator used for choosing numbers.
    protected Random m_randomGenerator = new Random();

    /**
     * Returns a random letter between A-Z, distributed the same as the frequency within the 
     * english language.
     */
    protected char randomLetter()
    {
        double num = m_randomGenerator.nextDouble();

        double sum = 0;
        for (int i = 0; i < LETTER_FREQUENCIES.length; i++)
        {
            sum += LETTER_FREQUENCIES[i];

            if (sum >= num) {
                return (char) (i+'A');
            }
        }
        return 'z';
    }

    /**
     * Generates a width x height baord with letters randomly distributed according
     * to the letter frequencies within the English language.
     */
    public Board generateRandomBoard(int width, int height)
    {
        Collection<Cell> cells = new ArrayList<Cell>();

        for (int row = 0; row < height; row++)
        {
            for(int col = 0; col < width; col++)
            {
                cells.add(new Cell(row, col, randomLetter()));
            }
        }
        
        Board board = new Board(width, height, cells);

        return board;
    }

}
