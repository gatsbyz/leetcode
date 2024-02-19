func numWays(n int, k int) int {
    // Handle base cases: no posts or no colors.
    if n == 0 || k == 0 {
        return 0
    }
    // If there's only one post, it can be painted with any of the k colors.
    if n == 1 {
        return k
    }
    
    // Initialize variables:
    // sameColor - ways to paint where the last two posts are of the same color.
    // diffColor - ways to paint where the last two posts are of different colors.
    var (
        sameColor = k       // For 2 posts, both can be the same color in k ways.
        diffColor = k * (k - 1) // For 2 posts, they can be different in k*(k-1) ways.
    )
    
    // Loop through posts from the third one to the nth post.
    for i := 2; i < n; i++ {
        // Temporary storage for diffColor to update sameColor in the next step.
        tmp := diffColor
        // Update diffColor for the next post:
        // Can choose any of the k-1 colors different from the last post's color.
        // The total ways include both previous sameColor and diffColor scenarios.
        diffColor = (sameColor + diffColor) * (k - 1)
        // Update sameColor for the next post: it equals the previous diffColor,
        // as we're now considering these configurations for the "same color" scenario.
        sameColor = tmp
    }
    
    // Return the total ways to paint the fence by summing sameColor and diffColor ways.
    // This includes all configurations up to the nth post.
    return sameColor + diffColor
}