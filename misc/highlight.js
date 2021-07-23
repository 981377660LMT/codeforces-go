// ==UserScript==
// @name         Codeforces Highlight
// @namespace    EndlessCheng
// @version      0.1
// @description  highlight some important words
// @author       EndlessCheng
// @match        https://atcoder.jp/contests/*
// @exclude      https://atcoder.jp/contests/*/submissions/*
// @match        https://codeforces.com/contest/*
// @match        https://codeforces.com/group/*
// @match        https://codeforces.com/gym/*
// @match        https://codeforces.com/problemset/problem/*
// @match        https://codingcompetitions.withgoogle.com/*
// @match        https://leetcode-cn.com/contest/*
// @match        https://www.acwing.com/problem/content/*
// ==/UserScript==

(function () {
    'use strict';

    // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions#Escaping
    function escapeRegExp(s) {
        return s.replace(/[.*+\-?^${}()|[\]\\]/g, '\\$&'); // $& means the whole matched string
    }

    String.prototype.replaceAll = function (oldSubstr, newSubstr) {
        return this.replace(new RegExp(escapeRegExp(oldSubstr), "g"), newSubstr);
    };

    String.prototype.highlight = function (substr, color) {
        return this.replaceAll(substr, "<span style='color: " + color + "'>" + substr + "</span>");
    };

    // 关键词高亮
    const color = "#49cc54"; // 49cc54  008dde
    // 高亮的词，一部分类似 Python 的关键字，另一部分是一些术语、修饰词之类
    // 注意若两个词有包含关系，把长的放前面
    const words = [
        // 描述
        "Initially", "initially", "guaranteed", "No matter", "no matter",
        "Not ", "not ", "don't", "didn't", "doesn't", "can't", "No ", " no ", "Non-", "non-", "without", "forbidden", "invalid", "nothing",
        "if and only if", "as long as",
        "Each ", " each ", // each 是个挺有趣的词，高亮它能帮助快速定位后面所描述的对象
        "every", " both ", " other ", " all ", "All ",
        "equal", "same", "different", "unique", "distinct", "strictly",
        "must", "Only", "only", "just", "Exactly", "exactly", "always", "indeed",
        "pairwise", "adjacent", "neighbour", "in a row", "consecutive", "contiguous", "one after another", "as possible",

        // 求解
        "minimize", "maximize", "minimum", "maximum", "minimal", "maximal", "smallest", "largest", "shortest", "longest",
        " small", " big", " large", " few",
        "At least", "At most", "at least", "at most",
        "more", "much",

        // 特殊描述
        "Empty", "empty",
        "zero", "positive", "negative",
        "decreasing", "descending", "increasing", "ascending", "sorted",
        "permutations", "permutation",
        "lowercase", "lower case", "uppercase", "upper case",
        "lexicographically", "lexicographical", "palindrome",
        "undirected", "directed", "bidirectional",
        "independently", "independent",
        "expected value",
        " circle", " ring",
        "counterclockwise", "counter-clockwise", "clockwise",
        "origin", "initial",
        "infinite",

        "没有", "不是", "不同", "不需要", "不会",
        "相同",
        "所有", "每个", "任何", "任意", "和", "与", "并且", "且", "同时", "或", "之一", "反之", "必须", "仅", "其他", // todo regex 每...
        "最小", "最大", "最少", "最多", "最短", "最长", "最早", "最晚", "最后", "第一", // todo regex 最...
        "至少", "至多", "恰好",
        "非空", "连续", "子数组", "子区间", "区间", "子序列", "子字符串", "子串",
        "严格", "递增", "递减", "升序", "降序", "字典序",
        "重复", "重新", "相邻",
        "叶子", "叶节点",
        "单向", "双向",
        "返回", "计算",
    ];

    const tags = ['p', 'li'];
    for (let ti = 0; ti < tags.length; ti++) {
        let pNodes = document.getElementsByTagName(tags[ti]);
        for (let i = 0; i < pNodes.length; i++) {
            let text = pNodes[i].innerHTML;

            for (let j = 0; j < words.length; j++) {
                text = text.highlight(words[j], color);
            }

            // 额外高亮
            const colorGreen = "#ff0000";
            text = text.highlight("取模", colorGreen).highlight("取余", colorGreen);

            text = text.replaceAll("Mr. ", "Mr.")
                .replaceAll("mr. ", "mr.")
                .replaceAll("Dr. ", "Dr.")
                .replaceAll("I.e. ", "I.e.")
                .replaceAll("i.e. ", "i.e.")
                .replaceAll("i. e. ", "i.e.")
                .replaceAll("i. e. ", "i.e.") // see https://codeforces.com/contest/1535/problem/A
                .replaceAll(". \\", ".\\") // 防止数学公式异常
                .replaceAll("...", "⋯") // 特殊处理一些句点，这些是不需要换行处理的
                // So you decided to hold a contest on Codeforces.
                // The maximum size of an array is $$$k$$$.
                .replaceAll(". ", ".</p><p>") // 加个换行（英文）
                .replaceAll(".\n", ".</p><p>") // 加个换行（英文）
                .replaceAll("。", "。</p><p>") // 加个换行（中文）
            // .replaceAll("\\dots", "~.~.~.~") // 替换掉省略号
            // .replaceAll("\\ldots", "~.~.~.~"); // 替换掉省略号

            // .replace(/(\$\$\$.+?\$\$\$)/g, "‘$1’"); // 教训：不应该加这个，看似优化实则是帮倒忙

            pNodes[i].innerHTML = text;
        }
    }
})();
