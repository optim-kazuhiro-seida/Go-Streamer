package sample

type (
    String struct {}
    Bool struct {}
    Float32 struct {}
    Float64 struct {}
    Integer struct {}
    Int32 struct{}
    Int64 struct{}
    sampleXXX struct{} // 先頭が小文字の場合はジェネレートされない
    Sample struct {
        Str string
        Int int
    }
    Sample0 struct {}
    Sample1 struct {}
    Sample2 struct {}
    Sample3 struct {}
)

type Sample4 struct {}
type Sample5 struct {
    aaa string
    bbbb string
}

