func designerPdfViewer(h []int32, word string) int32 {
    n := 1
    for _, v := range word {
        if int(h[int(v)-97]) > n {
            n = int(h[int(v)-97])
        }
    }
    return int32(n*len(word))
}
