package com.zirain.example;

import java.util.List;

import org.apache.commons.lang3.tuple.ImmutablePair;
import  org.apache.commons.lang3.tuple.Pair;

public class FibonacciPalindromeImpl implements FibonacciPalindrome {
    public Pair<Integer, Integer> find(List<Integer> seq) {
        Integer len = seq.size();
        boolean[][] dp = new boolean[len][len];
        Integer ansStartIndex = -1;
        Integer ansLen = -1;
        for (Integer i = 0; i < len; i++) {
            for (Integer j = 0; j <= i; j++) {
                Integer curLen = i - j + 1;
                if (curLen < 3) {
                    dp[j][i] = (seq.get(i) == seq.get(j));
                } else if (curLen == 3) {
                    dp[j][i] = dp[j + 1][i - 1] && (seq.get(i) == seq.get(j));
                } else {
                    dp[j][i] = dp[j + 1][i - 1] && (seq.get(i) == seq.get(j))
                            && ((seq.get(j) + seq.get(j + 1)) == seq.get(j + 2)
                                || seq.get(j) == seq.get(j + 1) + seq.get(j + 2)
                                || seq.get(j) == seq.get(j+2));
                }

                if (dp[j][i] && (curLen > ansLen)) {
                    ansStartIndex = j;
                    ansLen = curLen;
                }
            }
        }

        return new ImmutablePair<>(ansStartIndex, ansLen);
    }
}
