package zhconv

import (
    "testing"
)


func BenchmarkH2z(b *testing.B) {
    input := "ｱｲｳｴｵ123ｶﾞｷﾞｸﾞｹﾞｺﾞﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟHello, World!"
    for i := 0; i < b.N; i++ {
        H2z(input)
    }
}

func TestH2z(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"", ""},
        {"ｱｲｳｴｵ", "アイウエオ"},
        {"123", "１２３"},
        {"ｶﾞｷﾞｸﾞｹﾞｺﾞ", "ガギグゲゴ"},
        {"ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ", "パピプペポ"},
        {"Hello, World!", "Ｈｅｌｌｏ，　Ｗｏｒｌｄ！"},
    }

    for _, test := range tests {
        result := H2z(test.input)
        if result != test.expected {
            t.Errorf("H2z(%q) = %q; want %q", test.input, result, test.expected)
        }
    }
}
