

document.addEventListener("DOMContentLoaded", function(){   
 
    document.getElementById("1").onclick = function() {set_device("1")};
    document.getElementById("2").onclick = function() {set_device("2")};
    document.getElementById("0").onclick = function() {set_device("0")};

    

 

/*

    if(next_button){
        
        document.getElementById("next").onclick = function(){next_page()};
        console.log(next_button.dataset.page)
    }
    if (previous_button){
        document.getElementById("previous").onclick = function(){previous_page()};
    }
    

    function next_page(){   
            
            var url = new URL(window.location.href);
            var search_params = url.searchParams;        
            search_params.set('page', next_button.dataset.page);
            url.search = search_params.toString();        
            var new_url = url.toString();      
            window.location = new_url;          
    }

    
    function previous_page(){             
     
            
        var url = new URL(window.location.href);
        var search_params = url.searchParams;        
        search_params.set('page', previous_button.dataset.page);
        url.search = search_params.toString();        
        var new_url = url.toString();           
       
        window.location = new_url;  
    
    

}
    */




    function set_device(device){    
        
        

        
            var url = new URL(window.location.href);
            var search_params = url.searchParams;        
            search_params.set('device', device);     
            url.search = search_params.toString();        
            var new_url = url.toString();           
            window.location = new_url;                       
    }
   
    
   

})




/*
function remove(){
    var url = new URL(window.location.href);
    var search_params = url.searchParams;        
    search_params.delete('sort');
    url.search = search_params.toString();        
    var new_url = url.toString();           
    window.location = new_url;  

}

*/
