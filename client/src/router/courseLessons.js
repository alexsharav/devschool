import pythonBg from "@/views/bgImages/courses-python-bg.png";
import dataAnalysisBg from "@/views/bgImages/courses-data-analysis-bg.png";

export default {
  Python: {
    title: "Основы программирования на Python",
    iLink: pythonBg,
    cLink: "/Python",
    cDescription:
      "Python – это мощный и простой в изучении язык, который отлично подходит для начинающих. На курсе вы с нуля освоите синтаксис Python, научитесь писать чистый и структурированный код, а также создадите свои первые программы.",
    rLayout: false,
    chapters: [
      {
        chapterTitle: "Основы Python",
        lessons: [
          { path: "/courses/Python/1/1.1", title: "Переменные и типы данных" },
          { path: "/courses/Python/1/1.2", title: "Условия и циклы" },
        ],
      },
      {
        chapterTitle: "Функции и модули",
        lessons: [
          { path: "/courses/Python/2/2.1", title: "Определение функций" },
        ],
      },
    ],
  },

  DataAnalysis: {
    title: "Основы анализа данных",
    iLink: dataAnalysisBg,
    cLink: "/DataAnalysis",
    cDescription:
      "Понятное и доступное введение в работу с данными. Вы узнаете, как анализировать данные, научитесь находить важную информацию в данных, строить наглядные выводы и делать первые шаги в сторону профессиональной аналитики. Этот курс поможет вам начать путь в сфере анализа данных.",
    rLayout: true,
    chapters: [
      {
        chapterNumber: 1,
        chapterTitle: "Введение в анализ данных",
        lessons: [
          { path: "/courses/DataAnalysis/1/1.1", title: "Что такое данные?" },
        ],
      },
    ],
  },
};
