package model;

/**
 * Represents a word.
 */
public class Word
{
    protected final String m_word;

    public Word(String word)
    {
        m_word = word;
    }

    public String getValue()
    {
        return m_word;
    }

    public boolean equals(Object that)
    {
        if (that == null || !(that instanceof Word))
            return false;

        Word thatWord = (Word) that;

        return m_word != null && m_word.equals(thatWord.getValue());
    }

}
