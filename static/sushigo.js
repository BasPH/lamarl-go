$(function () {
    $('#submit-simulate').on('click', function (e) {
        e.preventDefault(); // disable the default form submit event

        var data = $('#simulate-form').serialize();
        $.ajax({
            url: '/simulate',
            type: 'POST',
            data: data,
            success: function (response) {
                $('#simulate-result').text(response);
            },
            error: function (response) {
                console.log("Error: " + response);
            },
        });
    });
});