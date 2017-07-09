

/** City autocomplete field */

function split(val) {
    return val.split(/,\s*/);
}

function extractLast(term) {
    return split(term).pop();
}

function extractFirst(term) {
    return split(term)[0];
}

jQuery(function () {
    var $citiesField = jQuery("#city");

    $citiesField.autocomplete({
        source: function (request, response) {
            jQuery.getJSON(
                "http://gd.geobytes.com/AutoCompleteCity?callback=?&q=" + extractLast(request.term),
                function (data) {
                    response(data);
                }
            );
        },
        minLength: 3,
        select: function (event, ui) {
            var selectedObj = ui.item;
            placeName = selectedObj.value;
            if (typeof placeName == "undefined") placeName = $citiesField.val();

            if (placeName) {
                var terms = split($citiesField.val());
                // remove the current input
                terms.pop();
                // add the selected item (city only)
                terms.push(extractFirst(placeName));
                // add placeholder to get the comma-and-space at the end
                terms.push("");
                $citiesField.val(terms.join(", "));
            }

            return false;
        },
        focus: function() {
            // prevent value inserted on focus
            return false;
        },
    });

    $citiesField.autocomplete("option", "delay", 100);
});
