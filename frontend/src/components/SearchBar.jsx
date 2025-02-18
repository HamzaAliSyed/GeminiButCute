import { TbMessageSearch } from "react-icons/tb";

const SearchBar = () => {
    return (
        <>
        <form className="w-full max-w-[480px] pt-3 pl-2 relative">
            <div className="relative">
                <input type="search" placeholder="Find Chat" className="w-full p-4 rounded-full bg-slate-600 border-none focus:outline-none placeholder:text-gray-400 text-white" />
                <button className="absolute right-1 top-1/2 p-4 -translate-y-1/2 rounded-full bg-slate-500 hover:bg-slate-400 transition-colors"> <TbMessageSearch /> </button>
            </div>
        </form>
        </>
    );
};

export default SearchBar;