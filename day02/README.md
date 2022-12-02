I started having a [conversation with ChatGPT](./chat_gpt_convo.md) about this
one. I wanted it to simplify the challenge initially, but it gave some
pseudocode. It turns out that if you're insistent, it'll give you solutions in
Go.

I took the part 1 proposed solution verbatim, you can see it [here][initial guess].
It needed [some tweaks][part 1 compare] to get it to give the correct answer,
mostly that it got confused about the format of the strategy.

By the time I'd got to part 2 I'd lost the tab with the chat thread, so I couldn't
use the context to carry on. I switched back to using Copilot to suggest the
changes. I found I needed to think a lot more about what I needed to do here,
changing the method signatures myself and variable names myself. Copilot is less
good at modifying code than writing code. [Here's my changes][part 2] for part 2.

[initial guess]: https://github.com/boffbowsh/aoc2022/commit/0dd3d8245644a12528604183ad086172b046a22b#diff-8349227e2dd886f7cbf0c2d6a6ec84b6071d128d4abd684a4b70df45a905baef
[part 1 compare]:https://github.com/boffbowsh/aoc2022/compare/boffbowsh:0dd3d82...boffbowsh:a88827e
[part 2]: https://github.com/boffbowsh/aoc2022/commit/f6e85b5798ed587d844e5d9c0b5ce9683d7b761d
