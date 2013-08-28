package model;

import java.util.HashMap;
import java.util.Map;

/**
 * A node within a Trie.
 *
 * @author jbowens
 */
public class TrieNode
{

    // The parent of this node
    protected TrieNode m_parent;

    // If this node marks the end of any word
    protected boolean m_isWord;

    /* We expect our tries for boards to be pretty sparse, so creating an
     * array for every node would likely waste memory. Instead, we have a
     * map of all children.
     */
    protected Map<Character,TrieNode> m_children;

    /* For backtracking, we also maintain a map from node instances to
     * characters.
     */
    protected Map<TrieNode,Character> m_backtrack;

    public TrieNode()
    {
        m_children = new HashMap<Character,TrieNode>();
        m_backtrack = new HashMap<TrieNode,Character>();
    }

    public boolean isWord()
    {
        return m_isWord;
    }

    public void setIsWord()
    {
        m_isWord = true;
    }

    public boolean hasParent()
    {
        return m_parent != null;
    }

    public TrieNode getParent()
    {
        return m_parent;
    }

    public String getAncestralString()
    {
        StringBuilder builder = new StringBuilder();

        TrieNode curr = this;
        while (curr.hasParent())
        {
            TrieNode parent = curr.getParent();
            char c = parent.getCharacterToChild(curr);
            builder.insert(c, 0);
            curr = parent;
        }

        return builder.toString();
    }

    public char getCharacterToChild(TrieNode child)
    {
        return m_backtrack.get(child);
    }

    /**
     * Retrieves the immediate child of the current node that contains the
     * given character, if it exists. May return null if no such node exists.
     *
     * @param c the character to look up
     */
    public TrieNode getChildForCharacter(char c)
    {
        Character character = new Character(c);
        return m_children.containsKey(character) ? m_children.get(character) : null;
    }

    /**
     * Retrieves the descendant of this node that contains the given substring,
     * if it exists. May return null if no such node exists.
     *
     * @param substr the string to lookup in the trie
     */
    public TrieNode getDescendantForString(String substr)
    {
        TrieNode curr = this;
        for (int i = 0; i < substr.length() && curr != null; i++)
        {
            char c = substr.charAt(i);
            curr = curr.getChildForCharacter(c);
        }

        return curr;
    }

    /**
     * Inserts the given node as a child of this node. Throws illegal argument exception if the node is indexed
     * by a character that already exists as a child of this node.
     */
    public void insertNode(char c, TrieNode newNode)
    {
        // Verify that this node isn't overwriting an existing node.
        TrieNode similarChild = getChildForCharacter(c);

        if (similarChild != null)
        {
            throw new IllegalArgumentException("A node already exists at this element with the character '"+c+"'");
        }

        m_children.put(c, newNode);
        m_backtrack.put(newNode, c);
        newNode.m_parent = this;
    }

    /**
     * Inserts the given string into the trie. Does nothing if the string already existed in
     * the trie.
     *
     * @return <code>true</code> if the string was already in the trie
     */
    public boolean insertString(String substr)
    {
        TrieNode currParent = this;
        for (int i = 0; i < substr.length(); i++)
        {
            TrieNode newNode = currParent.getChildForCharacter(substr.charAt(i));
            if (newNode == null)
            {
                newNode = new TrieNode();
                currParent.insertNode(substr.charAt(i), newNode);
            }
            currParent = newNode;
        }

        boolean existed = currParent.isWord();

        // Mark the final node as a word
        currParent.setIsWord();

        return existed;
    }

}
