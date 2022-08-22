#include <iostream>
#include <iterator>
#include <string>
#include <regex>

int main()
{
    std::string stat_name = "cluster.xds-grpc.upstream_rq_200";

    std::regex self_regex("REGULAR EXPRESSIONS",
                          std::regex_constants::ECMAScript | std::regex_constants::icase);
    if (std::regex_search(stat_name, self_regex))
    {
        std::cout << "Text contains the phrase 'regular expressions'\n";
    }
    
    std::vector<std::string> regexes;
    regexes.push_back("(response_code=\\.=(.+?);\\.;)|_rq(_(\\d{3}))");
    regexes.push_back("_rq(_(\\d{3}))|(response_code=\\.=(.+?);\\.;)");

    std::cout << "run regex search for " << stat_name << std::endl;
    for(std::string reg : regexes)//size()容器中实际数据个数 
    {
        std::cout << "regex: " << reg << std::endl;
        std::regex word_regex(reg);
        std::match_results<std::string::iterator> match;
        if (std::regex_search<std::string::iterator>(stat_name.begin(), stat_name.end(), match, word_regex) && match.size() > 1)
        {
            uint idx;
            for(idx=1;idx<match.size();idx++){
                if (match[idx].str() != ""){
                    break;
                }
            }
            const auto& remove_subexpr = match[idx];

            const auto& value_subexpr = idx+1 < match.size() ? match[idx+1] : remove_subexpr;

            // Determines which characters to remove from stat_name to elide remove_subexpr.
            std::string::size_type start = remove_subexpr.first - stat_name.begin();
            std::string::size_type end = remove_subexpr.second - stat_name.begin();
            std::cout << "remove_str: " << remove_subexpr.str() << "\n"
                    << "value_str: " << value_subexpr.str() << "\n"
                    << "start: " << start << "\n"
                    << "end: " << end << std::endl;
        }
    }
}