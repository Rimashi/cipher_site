$(document).ready(function () {
    // бургер и меню
    $(".burger").on("click", function () {
        $(this).toggleClass("active");
        // Изменяем aria-атрибут для доступности
        const isExpanded = $(this).hasClass('active');
        $(this).attr('aria-expanded', isExpanded);
        $(".side-menu, .overlay").toggleClass("show");
    });

    // закрытие по overlay
    $(".overlay").on("click", function () {
        $(".burger").removeClass("active").attr('aria-expanded', 'false');
        $(".side-menu, .overlay").removeClass("show");
    });

    // переключение вкладок
    $(".tab").on("click", function () {
        if ($(this).hasClass('active')) return; // Не делать ничего, если вкладка уже активна

        $(".tab").removeClass("active");
        $(this).addClass("active");
        let target = $(this).data("tab");

        // Плавное переключение контента
        $(".tab-pane.active").fadeOut(200, function () {
            $(this).removeClass("active");
            $("#" + target).fadeIn(200).addClass("active");
        });
    });

    // название из URL
    let params = new URLSearchParams(window.location.search);
    if (params.has("name")) {
        $(".cipher-title").text("Шифр " + params.get("name"));
    }

    // Закрытие меню при нажатии на Escape
    $(document).on('keydown', function (e) {
        if (e.key === "Escape") {
            $(".burger").removeClass("active").attr('aria-expanded', 'false');
            $(".side-menu, .overlay").removeClass("show");
        }
    });
});

// Обработка кнопок шифрования/дешифрования
$("button").on("click", function () {
    const $button = $(this);
    const $pane = $button.closest(".tab-pane");
    const action = $pane.attr("id");
    const cipherId = window.location.pathname.split("/").pop();

    const $textarea = $pane.find("textarea");
    const $keyInput = $pane.find(".key-input");
    const $result = $pane.find(".result");

    const text = $textarea.val().trim();
    const key = $keyInput.length ? $keyInput.val().trim() : "";

    if (!text) {
        $result.addClass("error").text("Пожалуйста, введите текст").addClass("show");
        return;
    }

    // Убираем класс ошибки если был
    $result.removeClass("error");

    // Показываем анимацию загрузки
    $result.addClass("show").html('<div class="loading">Обработка...</div>');

    // Для шифра Виженера получаем выбранный режим
    const data = { text, key };
    const $modeInput = $pane.find('input[name="mode"]:checked');
    if ($modeInput.length) {
        data.mode = $modeInput.val();
    }

    // Отправляем запрос на сервер
    $.ajax({
        url: `/api/${cipherId}/${action}`,
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (data) {
            $result.html(`<div class="result-content">${data.result}</div>`);
        },
        error: function (xhr) {
            const error = xhr.responseJSON && xhr.responseJSON.error ? xhr.responseJSON.error : "Произошла ошибка";
            $result.addClass("error").html(`<div class="result-content">${error}</div>`);
        }
    });
});