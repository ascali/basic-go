// console.log("home")
function LoadTypeMenu() {
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
    }
    
    fetch(`${BaseUrl}/read_type_menu`, requestOptions)
    .then(response => response.json())
    .then((result) => {
        let TypeMenuData = result.data
        let IdFoodMenu = document.querySelector("#id-food-menu")
        let IdFoodMenuData = document.querySelector("#id-food-menu-data")
        if (result.status==1) {
            if (TypeMenuData.length>0) {
                let htmlFoodMenu = `<li><a href="#" class="active" data-select="*"><b>Semua Menu</b></a></li>`
                TypeMenuData.map((value) => {
                    htmlFoodMenu += `<li><a href="#" data-select="data-type-${value.id}"><b>${value.name}</b></a></li>`
                })
                IdFoodMenu.innerHTML = htmlFoodMenu
    
                let htmlFoodMenuData = ""
                TypeMenuData.map((value) => {
                    htmlFoodMenuData += `
                        <div class="col-md-6 food-menu data-type-${value.id}">
                            <div class="sided-90x mv-30">
                            <div class="s-left"><img src="/static/images/chickenloverz.png" alt="chickenloverz"></div>
                            <div class="s-right">
                                <h5 class="mb-10"><b>Lengko</b><b class="color-primary float-right">Rp. 10.000,00</b></h5>
                                <p class="pr-70">Bumbu Lengko (tanpa nasi) dengan sayur segar, tempe, tahu dan sambal kacang.</p>
                                <a href="#" class="btn-brdr-primary plr-25"><b>Pesan</b></a>
                            </div>
                            </div>
                        </div>
                    `
                })
                IdFoodMenuData.innerHTML = htmlFoodMenuData
            } else {
                IdFoodMenu.innerHTML = `<li><a href="#" class="active" data-select="*"><b>Menu Not Found</b></a></li>`
                IdFoodMenuData.innerHTML = `
                    <div class="col-md-6 food-menu *">
                        <div class="sided-90x mv-30">
                        <div class="s-left"><img src="/static/images/chickenloverz.png" alt="chickenloverz"></div>
                        <div class="s-right">
                            <h5 class="mb-10"><b>Not Found</b><b class="color-primary float-right">Rp. xx.xxx,xx</b></h5>
                            <p class="pr-70">Description Not Found</p>
                            <a href="#" class="btn-brdr-primary plr-25"><b>Can't Order</b></a>
                        </div>
                        </div>
                    </div>
                `
            }
        }
    })
    .catch(error => console.log('error', error))
}
LoadTypeMenu()
