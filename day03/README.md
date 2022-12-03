The initial suggestion from ChatGPT didn't work. It misunderstood the priority
additions — it wanted to add the priorities of any type that appeared in _all_
rucksacks, giving an output of 0.

After fixing that, I noticed the rules suggest that each duplicated item gets
counted once per rucksack, whereas GPT counted each occurance. I clarified:

> Each duplicated item type should only be counted once per rucksack, rather than each duplicated one. What do I need to change?

It returned the same result. I told it the answer was too low. Its reply was incredible:

> It looks like the original solution was correct and the updated solution is not. The original solution counts the number of times each character appears in both compartments of each rucksack, and if a character appears in both compartments of every rucksack, it adds its priority to the sum. This is the correct way to solve the puzzle, because the puzzle description says that "the only item type that appears in both compartments is lowercase p" (for the first rucksack), "the only item type that appears in both compartments is uppercase L" (for the second rucksack), etc.
>
> The updated solution tries to only count a character once per rucksack, but this is not necessary and doesn't match the description of the puzzle. In fact, it is incorrect, because it fails to count characters that appear multiple times in both compartments of the same rucksack.
>
> Here is an example that shows how the updated solution is incorrect:
>
```go
rucksacks := []string{
  "ab", // first compartment: "a", second compartment: "b"
  "bac", // first compartment: "ba", second compartment: "c"
}
```
> With the original solution, charCount['a'] will be 2 because 'a' appears in both compartments of both rucksacks. With the updated solution, charCount['a'] will be 0 because 'a' is counted only once in the first rucksack, but not counted at all in the second rucksack.
> To fix the updated solution, you can simply remove the counted map and the counted[char] = true line, so that the solution behaves like the original one.

I replied:

> Both solutions returned the same result with my input data. Would it help if I sent the data I'm working with?

Its reply:

> Yes, that would be helpful. Can you please provide the data you are using, and also explain why you think the result is incorrect? That would help me understand what might be going wrong and how to help you better.
> Also, please note that the original solution is correct and the updated solution is not. The updated solution tries to only count a character once per rucksack, but this is not necessary and doesn't match the description of the puzzle. In fact, it is incorrect, because it fails to count characters that appear multiple times in both compartments of the same rucksack. If you are using the updated solution, you should switch back to the original one.

I provide the input data. It replies:

> I'm sorry, but I'm not able to run your code because I'm a large language model trained by OpenAI and I don't have access to external resources like the input data you provided. I can only provide general guidance based on the information you've given me.
>
> Based on the input data you provided, it looks like your code is not splitting the rucksacks into their two compartments correctly. In the original solution, this is done using the following lines of code:

```go
// Split the rucksack into its two compartments
first := rucksack[:len(rucksack)/2]
second := rucksack[len(rucksack)/2:]
```

> This uses the length of the rucksack string to determine where to split it into the first and second compartments. The first compartment contains the first half of the characters, and the second compartment contains the second half of the characters.
>
> For example, suppose the rucksack string is "abcdef". The length of this string is 6, so the first compartment contains the first 6/2 = 3 characters ("abc") and the second compartment contains the remaining 3 characters ("def").
>
> If you are not using this method to split the rucksacks into their two compartments, then the result of your code will be different from the result of the original solution. I suggest you check how you are splitting the rucksacks and make sure you are using the method described above.
>
> Please let me know if you have any other questions or if you need further clarification on anything.

I should stress that this thing is responding _instantly_. The speed is the only way I know this isn't a human I'm talking to.

It seems to have gone as far as it can though, so I'm back to human debugging for the first time.

The eventual solution was just where it had misunderstood the need use the count of item types, not just the priorities. When I did first tweak from the suggested code, I didn't look closely enough. It was ignoring the count already, I needed to start using it.

## Part 2

I tried feeding it part 2 of the puzzle, but the reply was truncated, even after 2 "Try Again"s.

I returned to Copilot instead. I mostly described functions by name and had Copilot fill them out, and used a single comment for fetching a single element from a set. I hoped it'd handle the error case where there was more than one element, but it didn't seem to bother.

## Thoughts

Surely the next evolution for Copilot has to be a conversational mode where it is more like pairing with someone. Engineers rarely write entirely new code, we're more often modifying existing code to add new behaviour or fix a bug. Given that ChatGPT seems to be capable of understanding the context of existing code and modifying it based on a conversation, we're not far away from giving VSCode a Jira ticket number and having it give us a first guess at an implementation of the ticket.

Tomorrow, depending on the puzzle, I'm going to try either getting ChatGPT to come up with a unit test first, or give it a unit test to satisfy.
