package usecases

import "testing"

func Test_SubstringUsecase_SearchMaxUniqueSubstring(t *testing.T) {
	//тестовые данные
	testTable := []struct {
		name string
		str  string
		res  string
	}{
		{
			name: "Пример из тз 1",
			str:  "abcabcbb",
			res:  "abc",
		},
		{
			name: "Пример из тз 2",
			str:  "bbbb",
			res:  "b",
		},
		{
			name: "Пример из тз 3",
			str:  "pwwkew",
			res:  "wke",
		},
		{
			name: "Все буквы уникальны",
			str:  "abcdefjh",
			res:  "abcdefjh",
		},
		{
			name: "Все буквы одинаковые",
			str:  "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			res:  "a",
		},
		{
			name: "Правильная последовательность в конце строки",
			str:  "aaaaaaaaaaaaaaaaaaab",
			res:  "ab",
		},
		{
			name: "Правильная последовательность в середине строки",
			str:  "sasasagfdasghjklzxcasassasaas",
			res:  "fdasghjklzxc",
		},
		{
			name: "Кирилица",
			str:  "ититититкуититит",
			res:  "итку",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			s := NewSubstring()
			// Тестирование функции
			res, _ := s.SearchMaxUniqueSubstring(testCase.str)
			// Сравнение результатов
			t.Log("Ожидалось  -", testCase.res)
			t.Log("Получилось -", res)
			if testCase.res != res {
				t.Errorf("Ошибка теста")
			}
		})
	}
}
