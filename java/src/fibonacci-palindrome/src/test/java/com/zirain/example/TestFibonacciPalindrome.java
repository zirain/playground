package com.zirain.example;

import java.util.Arrays;
import java.util.List;

import org.apache.commons.lang3.tuple.ImmutablePair;
import org.apache.commons.lang3.tuple.Pair;
import org.junit.Test;
import org.junit.Assert;

public class TestFibonacciPalindrome {

    @Test
    public void testFindFibonacciPalindrome() throws Exception {

        List<Integer> sequence = Arrays.asList(1);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<>(0, 1), result);
    }

    @Test
    public void testFindFibonacciPalindrome2() throws Exception {

        List<Integer> sequence = Arrays.asList(1, 2);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(0, 1), result);
    }

    @Test
    public void testFindFibonacciPalindrome3() throws Exception {

        List<Integer> sequence = Arrays.asList(1, 2, 3);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(0, 1), result);
    }

    @Test
    public void testFindFibonacciPalindrome4() throws Exception {

        List<Integer> sequence = Arrays.asList(1, 2, 1);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(0, 3), result);
    }

    @Test
    public void testFindFibonacciPalindrome5() throws Exception {

        List<Integer> sequence = Arrays.asList(1, 2, 2, 1);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(1, 2), result);
    }

    @Test
    public void testFindFibonacciPalindrome6() throws Exception {

        List<Integer> sequence = Arrays.asList(1, 2, 3, 2, 1);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(0, 5), result);
    }

    @Test
    public void testFindFibonacciPalindrome7() throws Exception {

        List<Integer> sequence = Arrays.asList(1, 2, 3, 3, 2, 1);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(2, 2), result);
    }

    @Test
    public void testFindFibonacciPalindrome8() throws Exception {

        List<Integer> sequence = Arrays.asList(3, 0, 3, 3, 0, 3);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(0, 6), result);
    }

    @Test
    public void testFindFibonacciPalindrome9() throws Exception {

        List<Integer> sequence = Arrays.asList(100, 101, 3, 0, 3, 3, 0, 3, 30, 20);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<>(2, 6), result);
    }

    @Test
    public void testFindFibonacciPalindrome10() throws Exception {

        List<Integer> sequence = Arrays.asList(5, 3, 9, 1, 10, 11, 10, 1, 9, 3, 5);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(2, 7), result);
    }

    @Test
    public void testFindFibonacciPalindrome12() throws Exception {

        List<Integer> sequence = Arrays.asList(5, 3, 11, 2, 9, 11, 9, 2, 11, 3, 5);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(2, 7), result);
    }

    @Test
    public void testFindFibonacciPalindrome13() throws Exception {

        List<Integer> sequence = Arrays.asList(5, 3, 11, 2, 9, 11, 9, 2, 11, 3, 5);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(2, 7), result);
    }

    @Test
    public void testFindFibonacciPalindrome14() throws Exception {

        List<Integer> sequence = Arrays.asList(5, 3, 10, 1, 10, 11, 10, 1, 10, 3, 5);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(2, 7), result);
    }

    @Test
    public void testFindFibonacciPalindrome15() throws Exception {

        List<Integer> sequence = Arrays.asList(1, 1, 1, 1, 1, 1, 1);
        FibonacciPalindromeImpl fibonacciPalindrome = new FibonacciPalindromeImpl();

        Pair<Integer, Integer> result = fibonacciPalindrome.find(sequence);
        Assert.assertEquals(new ImmutablePair<Integer, Integer>(0, 7), result);
    }
}
