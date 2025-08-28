document.addEventListener('DOMContentLoaded', function () {
    const cardNumberInput = document.getElementById('cardNumber');
    const cardHolderInput = document.getElementById('cardHolder');
    const expiryDateInput = document.getElementById('expiryDate');


    // Kart numarası formatlama
    cardNumberInput.addEventListener('input', function (e) {
        let value = e.target.value.replace(/\s+/g, '').replace(/[^0-9]/gi, '');
        let formattedValue = '';

        for (let i = 0; i < value.length; i++) {
            if (i > 0 && i % 4 === 0) {
                formattedValue += ' ';
            }
            formattedValue += value[i];
        }

        e.target.value = formattedValue;
    });

    // Son kullanma tarihi formatlama
    expiryDateInput.addEventListener('input', function (e) {
        let value = e.target.value.replace(/\D/g, '');

        if (value.length > 0) {
            value = value.match(new RegExp('.{1,2}', 'g')).join('/');
            if (value.length > 5) {
                value = value.substring(0, 5);
            }
        }

        e.target.value = value;
    });

    // CVV sadece sayı kabul etme
    document.getElementById('cvv').addEventListener('input', function (e) {
        e.target.value = e.target.value.replace(/\D/g, '');
    });

    // Kart sahibi sadece harf kabul etme
    cardHolderInput.addEventListener('input', function (e) {
        e.target.value = e.target.value.replace(/[^a-zA-ZğüşıöçĞÜŞİÖÇ\s]/g, '');
    });
});