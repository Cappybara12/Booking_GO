function Prompt() {
    const toast = function (c) {
        const {
            msg = "",
            icon = "success",
            position = "top-end",
        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.onmouseenter = Swal.stopTimer;
                toast.onmouseleave = Swal.resumeTimer;
            }
        });
        Toast.fire({});
    };

    const success = function (c) {
        const {
            msg = "",
            title = "",
            footer = ""
        } = c;
        Swal.fire({
            icon: "success",
            title: title,
            text: msg,
            footer: footer
        });
    };

    const error = function (c) {
        const {
            msg = "",
            title = "",
            footer = ""
        } = c;
        Swal.fire({
            icon: "error",
            title: title,
            text: msg,
            footer: footer
        });
    };

    const custom = async function (c) {
        const {
            icon="",
            msg = "",
            title = "",
            showConfirmButton:true,
        } = c;

        const { value: result } = await Swal.fire({
            icon:icon,
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            showConfirmButton:showConfirmButton,
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ];
            }
        });

      if (result){
        if(result.dismiss !== Swal.DismissReason.cancel){
            if(result !==""){
                if (c.callback !==undefined){
                    c.callback(result)
                }
            }else{
                c.callback(false);
            }
        }else{
            c.callback(false)
        }
      }  
    };

    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom
    };
}